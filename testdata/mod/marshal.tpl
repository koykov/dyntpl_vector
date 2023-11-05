{% ctx data = source|vector::parseJSON().(vector) %}
root:{%= data|vector::marshal() %};
nested0:{%= vector::marshal(data.x) %};
nested1:{%= data.x.y|vector::marshal() %};
nested2:{%= data.x.y.z|vector::marshal() %}.
