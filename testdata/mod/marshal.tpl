{% ctx data = source|vector::parseJSON().(vector) %}
root:{%= data|vector::marshal() %};
nested0:{%= data.x|vector::marshal() %};
nested1:{%= data.x.y|vector::marshal() %};
nested2:{%= data.x.y.z|vector::marshal() %}.
