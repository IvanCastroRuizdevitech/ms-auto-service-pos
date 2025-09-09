CREATE OR REPLACE FUNCTION public.fnc_get_eds_list()
 RETURNS json
 LANGUAGE plpgsql
AS $function$
/*************************************************************************************
    Archivo:        fnc_get_eds_list
    Autor:          Jose Rodriguez
    Fecha:          03-09-2025
    Descripción:    Obtiene la lista de EDS activas definidas en la tabla
                    sat.tbl_sat_datos_empresas.
*************************************************************************************/
DECLARE
    v_response JSON;
    v_error    TEXT;
BEGIN
    -- Construye un arreglo JSON a partir de la consulta final
    SELECT COALESCE(array_to_json(array_agg(row_to_json(t))), '[]')::json
    INTO v_response
    FROM (
        SELECT
            emp.id AS "value",
            emp.razon_social AS "label"
        FROM
            public.ct_empresas AS emp
        INNER JOIN
            sat.tbl_sat_datos_empresas AS sat_emp ON emp.id = sat_emp.ct_empresas_id
        WHERE
            emp.estado = 'A'
    ) t;

    -- Envuelve la data en la estructura de respuesta final
    v_response := json_build_object(
        'status', 200,
        'success', true,
        'data', v_response
    );

    RETURN v_response;

EXCEPTION
    WHEN OTHERS THEN
        GET STACKED DIAGNOSTICS v_error = MESSAGE_TEXT;
        v_response := json_build_object(
            'status', 500,
            'success', false,
            'message', 'Error interno en la base de datos.',
            'error_detail', v_error
        );
    RETURN v_response;
END;
$function$;
SELECT public.fnc_get_eds_list();