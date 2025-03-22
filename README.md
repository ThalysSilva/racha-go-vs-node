# Codigo utilizado no post do [linkedin](https://www.linkedin.com/feed/update/urn:li:activity:7308145428812750849/) para comparar velocidade de um fluxo de criação de pedido fictício entre go e node.

## Conteúdo do post:

Fiz um "RACHA" entre Go e Node. Cada um processando uma pipeline de 100 pedidos fictícios pra tirar a teima a velocidade dos dois.


O teste? Leitura de dados (200ms por pedido), cálculo de preços e salvamento no banco (300ms cada).


No node, como ele é um event loop single Thread, o fluxo foi sequencial durando aprox. 50s (20s e leitura + 30s de salvamento)

Já no Go, Utilizei as goroutines e canais, oque fez cada pipeline rodar em paralelo concorrentemente, aproveitando os múltiplos núcleos. Resultado aproximado ? 30 SEGUNDOS!! Ganho de 39,5% na execução!!!


O que Trouxe essa vantagem pro go foi na concorrência nativa dele: goroutines são leves e escalam bem, enquanto o Node, mesmo com os Workers (que são threads reais e mais pesadas), exigem mais esforço e overhead para chegar perto disso. Claro, o sucesso do node é no I/O leve e em prototipagens rápidas (MVP), mas, para sistemas que pedem performance e escala como esse "duelo" que eu fiz, sem duvidas o Go leva vantagem. E ai alguém já fez algo parecido ? Oque vocês acharam desse racha ???

