-- DROP FUNCTION public.fnc_obtener_configuracion_pos_autoservicio(jsonb);

CREATE OR REPLACE FUNCTION public.fnc_obtener_configuracion_pos_autoservicio(p_json_params jsonb)
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
  v_response             jsonb;

  v_surtidores           jsonb;
  v_surtidores_detalles  jsonb;
  v_empresas             jsonb;
  v_medios_pagos         jsonb;
  v_caras                int[];
  v_msg                  text;
  v_parametros           jsonb;
BEGIN
  -- Log NO bloqueante
  BEGIN
    INSERT INTO public.logs(texto_log, fecha_log, nombre_up, id_proceso)
    VALUES (p_json_params::text, NOW(), 'fnc_obtener_configuracion_pos_autoservicio', 0);
  EXCEPTION WHEN OTHERS THEN
    NULL;
  END;

  IF p_json_params ? 'caras' AND COALESCE(p_json_params->>'caras','') <> '' THEN
    v_caras := ARRAY(
      SELECT NULLIF(trim(x),'')::int
      FROM unnest(string_to_array(p_json_params->>'caras', ',')) AS t(x)
      WHERE trim(x) ~ '^\d+$'
    );
  END IF;

  /* =========================
     Colecciones base (ordenado)
     ========================= */

  -- surtidores
  SELECT COALESCE(jsonb_agg(to_jsonb(s)), '[]'::jsonb)
    INTO v_surtidores
  FROM (
    SELECT *
    FROM public.surtidores
    ORDER BY id
  ) s;


  -- empresas
  SELECT COALESCE(jsonb_agg(to_jsonb(e)), '[]'::jsonb)
    INTO v_empresas
  FROM (
    SELECT *
    FROM public.empresas
    ORDER BY id
  ) e;

 
   /* ============================================
     medios_pagos con mp_atributos.autoservicio = true
     ============================================ */
  SELECT COALESCE(jsonb_agg(to_jsonb(mp)), '[]'::jsonb)
    INTO v_medios_pagos
  FROM (
    SELECT *
    FROM public.medios_pagos
    WHERE COALESCE( (mp_atributos->>'autoservicio')::boolean, false ) IS TRUE
    -- Si también quieres solo activos: AND estado = 'A'
    ORDER BY id
  ) mp;

  /* ============================================
     surtidores_detalles filtrado por caras y JOIN
     ============================================ */

  /*
    - Se filtra por sd.cara IN v_caras (si viene el parámetro).
    - JOIN productos (p) para obtener descripción/precio.
    - LEFT JOIN productos_familias (pf) para familia/código/atributos.
    - LEFT JOIN LATERAL para traer el ÚLTIMO precio de familia (fp_last).
    - Orden determinístico por cara, manguera, id.
  */
  SELECT COALESCE(jsonb_agg(row_json), '[]'::jsonb)
    INTO v_surtidores_detalles
  FROM (
    SELECT
      to_jsonb(sd)
      || jsonb_build_object(
           'producto_descripcion', p.descripcion,
           'producto_precio',      p.precio,
           'familia_id',           pf.id,
           'familia_codigo',       pf.codigo,
           'familia_atributos',    pf.atributos,
           'precio_familia',       fp_last.precio
         ) AS row_json
    FROM public.surtidores_detalles sd
    JOIN public.productos p
      ON p.id = sd.productos_id
    LEFT JOIN public.productos_familias pf
      ON pf.id = p.familias
    LEFT JOIN LATERAL (
      SELECT fp.precio
      FROM public.familia_precios fp
      WHERE fp.familia_id::bigint = pf.id
      ORDER BY fp.fecha_actualizado DESC NULLS LAST, fp.id DESC
      LIMIT 1
    ) AS fp_last ON TRUE
    WHERE (v_caras IS NULL OR sd.cara = ANY (v_caras))
    ORDER BY sd.cara, sd.manguera, sd.id
  ) q;
 
   /* ============================================
     wacher_parametros: POS_ID y POS_PRINCIPAL
     ============================================ */
  SELECT COALESCE(
           jsonb_object_agg(wp.codigo, jsonb_build_object('valor', wp.valor, 'tipo', wp.tipo)),
           '{}'::jsonb
         )
    INTO v_parametros
  FROM public.wacher_parametros wp
  WHERE wp.codigo IN ('POS_ID', 'POS_PRINCIPAL');

  /* =========================
     Respuesta final
     ========================= */
  v_response := jsonb_build_object(
    'status',  200,
    'mensaje', 'Configuración obtenida con éxito',
    'data', jsonb_build_object(
      'surtidores',          COALESCE(v_surtidores, '[]'::jsonb),
      'surtidores_detalles', COALESCE(v_surtidores_detalles, '[]'::jsonb),
      'empresas',            COALESCE(v_empresas, '[]'::jsonb),
      'medios_pagos',        COALESCE(v_medios_pagos, '[]'::jsonb),
      'parametros',          COALESCE(v_parametros, '{}'::jsonb)
    )
  );

  RETURN v_response;

EXCEPTION WHEN OTHERS THEN
  GET STACKED DIAGNOSTICS v_msg = MESSAGE_TEXT;
  RETURN jsonb_build_object(
    'status', 500,
    'mensaje', 'Error al obtener la configuración: ' || COALESCE(v_msg, 'desconocido')
  );
END;
$function$;
