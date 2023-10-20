{% ctx data = source|vector::parseYAML().(vector) %}
{%= data.x.y.z %}
