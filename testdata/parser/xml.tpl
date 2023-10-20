{% ctx data = source|vector::parseXML().(vector) %}
{%= data.x.y.z %}
