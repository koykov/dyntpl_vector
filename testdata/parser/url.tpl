{% ctx data = source|vector::parseURL().(vector) %}
{%= data.query.xyz %}
