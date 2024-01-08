# Arquitectura Hexagonal
La arquitectura hexagonal, también conocida como "Ports and Adapters", es un patrón de diseño de software donde la lógica de negocio (o dominio) es independiente a la lógica del usuario y la lógica del servidor. 
Los componentes son conectados entre sí a través de `Puertos` que podemos encontrar en el codigo como interfaces, y los `Adaptadores` deben ser implementaciones de estos puertos que permitiran la comunicacion de componentes externos con el dominio o bien, la comunicacion del dominio con componentes externos.

![Hexagonal Architecture](hexagonal.png)

La arquitectura hexagonal se compone por los siguientes elementos:

- Dominio: En el centro de la arquitectura hexagonal está la lógica del negocio. Este es el corazón de la aplicación y contiene reglas, algoritmos y procesos de negocio. Aqui se contienen los casos de uso, que son una descripción abstracta de lo que hacen los usuarios con nuestra aplicación, y se contienen los puertos por los que se podrán comunicar con el dominio y el dominio con componentes externos.
- Puertos (Ports): Los puertos son interfaces que definen cómo se pueden comunicar con el dominio mediante el principio de inversión de dependencias.
    - Puerto de Entrada: Es una interfaz simple que componentes externos pueden llamar para conectar con los Casos de Uso. Estos componentes los encontraremos representados graficamente a la izquierda del hexagono.
    - Puerto de Salida: Es una interfaz simple a la que nuestros Casos de Uso llaman si necesitan algo de componentes externos. Caso contrario, estos los encontraremos representados graficamente del lado derecho del hexagono.
- Adaptadores: Los adaptadores son implementaciones de los puertos. Al llamarlos conectamos la lógica del negocio con el mundo exterior.







