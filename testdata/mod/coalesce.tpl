{% ctx data = source|vector::parseJSON().(vector) %}
{%= data|vector::coalesce("x.y.z.a.b.c", "x.y.z.a.b", "x.y.z.a", "x.y.z") %}
