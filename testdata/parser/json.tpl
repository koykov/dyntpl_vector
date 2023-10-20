{% ctx data = source|vector::parseJSON().(vector) %}
{%= data.x.y.z %}
