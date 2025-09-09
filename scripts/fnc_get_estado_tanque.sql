CREATE OR REPLACE FUNCTION public.fnc_get_estados_tanque()
 RETURNS json
 LANGUAGE plpgsql
AS $function$
/*************************************************************************************
    Archivo:        fnc_get_estados_tanque
    Autor:          Jose Rodriguez
    Fecha:          03-09-2025
    Descripción:    Lee la tabla tbl_estados y la transforma al formato 
                    requerido por la HU 036 (O/F).
*************************************************************************************/
DECLARE
    v_response JSON;
BEGIN
    -- Construimos el JSON a partir de la tabla tbl_estados
    SELECT COALESCE(array_to_json(array_agg(row_to_json(t))), '[]')::json
    INTO v_response
    FROM (
        SELECT
           id_estado AS "id",
            CASE
                WHEN id_estado = 1 THEN 'O'  -- Si es 1 (ACTIVO), la clave es 'O'
                ELSE 'F'                      -- Si no (INACTIVO), la clave es 'F'
            END AS "clave",
            CASE
                WHEN id_estado = 1 THEN 'En Operación'
                ELSE 'Fuera de Operación'
            END AS "descripcion"
        FROM public.tbl_estados
        WHERE id_estado IN (1, 2)
    ) t;

    -- Envolvemos en la respuesta estándar
    v_response := json_build_object(
        'status', 200,
        'success', true,
        'data', v_response
    );

    RETURN v_response;
END;
$function$;
select public.fnc_get_estados_tanque();
