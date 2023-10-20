{% ctx data = source|vector::parseHAL().(vector) %}
{%= data.0.code %}-{%= data.0.script %}-{%= data.0.region %};q={%= data.0.quality %}
