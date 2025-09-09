-- DROP FUNCTION public.fnc_consultar_catalogos_tanques(varchar, int8);

CREATE OR REPLACE FUNCTION public.fnc_consultar_catalogos_tanques(catalogo_nombre character varying, filtro_id bigint DEFAULT NULL::bigint)
 RETURNS json
 LANGUAGE plpgsql
AS $function$
/*************************************************************************************
    Archivo:        fnc_consultar_catalogos_tanques (Versión Definitiva)
    Autor:          Jose Rodriguez
    Fecha:          04-09-2025
    Descripción:    Función centralizada para obtener los catálogos de familias y 
                    productos para la HU de Creación de Tanques.
*************************************************************************************/
DECLARE
    v_response JSON;
    v_target_product_type_id BIGINT;
BEGIN
    -- Caso 1: El frontend pide la lista de FAMILIAS
    IF catalogo_nombre = 'FAMILIAS' THEN
        SELECT COALESCE(array_to_json(array_agg(row_to_json(t))), '[]')::json
        INTO v_response
        FROM (
            SELECT id_tipo_producto AS value, descripcion AS label 
            FROM public.tbl_tipo_productos 
            WHERE id_tipo_producto IN (3, 4) -- IDs para GASOLINA y DIESEL
        ) t;

    -- Caso 2: El frontend pide la lista de PRODUCTOS para una familia específica
    ELSIF catalogo_nombre = 'PRODUCTOS' THEN
        IF filtro_id IS NULL THEN
            RAISE EXCEPTION 'Para consultar productos, se requiere un filtro_id.';
        END IF;

        -- "Traduce" el ID de la familia (3 o 4) al tipo de producto (28 o 29)
        CASE filtro_id
            WHEN 3 THEN v_target_product_type_id := 28; -- GASOLINA -> tipo 28
            WHEN 4 THEN v_target_product_type_id := 29; -- DIESEL -> tipo 29
            ELSE v_target_product_type_id := -1; -- Un valor que no encontrará nada
        END CASE;

        -- Ejecuta la consulta de productos con el tipo ya traducido
        SELECT COALESCE(array_to_json(array_agg(row_to_json(t))), '[]')::json
        INTO v_response
        FROM (
            SELECT
                p.id AS "value",
                p.descripcion AS "label",
                COALESCE(scu.descripcion, 'N/A') AS "unit"
            FROM
                public.ct_productos AS p
            LEFT JOIN
                sat.tbl_asociacion_clave_sat_unidades AS acsu ON p.id = acsu.producto_id
            LEFT JOIN
                sat.tbl_sat_claves_unidades AS scu ON acsu.sat_clave_unidad_id = scu.id
            WHERE
                p.tipo_producto_id = v_target_product_type_id AND p.estado = 'A'
        ) t;
    END IF;

    -- Construye la respuesta final exitosa
    RETURN json_build_object(
        'status', 200,
        'success', true,
        'data', v_response
    );

EXCEPTION
    WHEN OTHERS THEN
        RETURN json_build_object(
            'status', 500,
            'success', false,
            'message', 'Error interno en la base de datos.',
            'error_detail', SQLERRM
        );
END;
$function$
;
