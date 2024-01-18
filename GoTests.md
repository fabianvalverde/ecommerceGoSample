
# Convenciones de Testing en Go

- ## Archivos de prueba: 
Estos deben tener el sufijo `_test.go`. Lo ideal es que el nombre comience con el mismo nombre que el archivo o estructura que se está probando.
- ## Funciones de prueba: 
En este proyecto se utiliza la estructura del paquete de pruebas de golang `testing.T` para la creacion y ejecucion de pruebas unitarias y de integración. Este tipo proporciona métodos para reportar fallos de pruebas y para la lógica de control de flujo de las pruebas. Cuando escribes una función de prueba en Go, pasas un puntero a testing.T (por ejemplo, func TestMyFunction(t *testing.T)) para manejar la ejecución de la prueba y reportar resultados. Por lo tanto, testing.T se asocia con pruebas unitarias y de integración. Por convencion, la función para este tipo de pruebas debe comenzar con "Test", seguida del nombre de la función que se está probando.

Adoptar convenciones propias para manejar datos que se repiten en distintas funciones es una excelente práctica dentro de un equipo de trabajo. En las pruebas de este proyecto, utilizaremos variables específicas para comparar los resultados obtenidos ('actual') con los deseados ('expected'). Esta convención nos ayudará a mantener la claridad y coherencia a lo largo de nuestras pruebas.

# ¿Qué es Mocking? 

En el contexto de testing es una técnica que permite simular el comportamiento de un componente concreto del sistema dentro de un conjunto de pruebas unitarias. Esta técnica es muy utilizada cuando el código que se desea probar interactúa con un componente externo, lo que nos permite simular el comportamiento de este obteniendo los resultados deseados. (Explicar donde vemos mocking)


