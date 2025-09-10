CREATE OR REPLACE FUNCTION public.fnc_obtener_configuracion_pos_autoservicio(p_json_params jsonb)
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
    v_response jsonb;
    v_surtidores jsonb;
    v_surtidores_detalles jsonb;
    v_productos jsonb;
    v_familias_precios jsonb;
    v_empresas jsonb;
    v_equipos jsonb;
BEGIN
    -- Obtener datos de surtidores
    SELECT jsonb_agg(row_to_json(s)) INTO v_surtidores FROM public.surtidores s;

    -- Obtener datos de surtidores_detalles
    SELECT jsonb_agg(row_to_json(sd)) INTO v_surtidores_detalles FROM public.surtidores_detalles sd;

    -- Obtener datos de productos
    SELECT jsonb_agg(row_to_json(p)) INTO v_productos FROM public.productos p;

    -- Obtener datos de familias_precios (asumiendo que es productos_familias)
    SELECT jsonb_agg(row_to_json(pf)) INTO v_familias_precios FROM public.productos_familias pf;

    -- Obtener datos de empresas
    SELECT jsonb_agg(row_to_json(e)) INTO v_empresas FROM public.empresas e;

    -- Obtener datos de equipos
    SELECT jsonb_agg(row_to_json(eq)) INTO v_equipos FROM public.equipos eq;

    -- Construir la respuesta JSON
    v_response := jsonb_build_object(
        'status', 200,
        'mensaje', 'Configuracion obtenida con exito',
        'data', jsonb_build_object(
            'surtidores', COALESCE(v_surtidores, '[]'::jsonb),
            'surtidores_detalles', COALESCE(v_surtidores_detalles, '[]'::jsonb),
            'productos', COALESCE(v_productos, '[]'::jsonb),
            'familias_precios', COALESCE(v_familias_precios, '[]'::jsonb),
            'empresas', COALESCE(v_empresas, '[]'::jsonb),
            'equipos', COALESCE(v_equipos, '[]'::jsonb)
        )
    );

    RETURN v_response;
EXCEPTION
    WHEN OTHERS THEN
        RETURN jsonb_build_object(
            'status', 500,
            'mensaje', 'Error al obtener la configuración: ' || SQLERRM
        );
END;
$function$;


