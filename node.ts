async function lerPedidos(): Promise<number[]> {
  const pedidos: number[] = [];
  for (let i = 1; i <= 5; i++) {
    await new Promise((resolve) => setTimeout(resolve, 200));
    pedidos.push(i);
  }
  return pedidos;
}

async function calcularPrecos(pedidos: number[]): Promise<number[]> {
  const precos: number[] = [];
  for (const pedido of pedidos) {
    precos.push(pedido * 10.5);
  }
  return precos;
}

async function salvarNoBanco(precos: number[]): Promise<string[]> {
  const resultados: string[] = [];
  for (const preco of precos) {
    await new Promise((resolve) => setTimeout(resolve, 300));
    resultados.push(`Salvo pedido com preço ${preco.toFixed(2)}`);
  }
  return resultados;
}

async function main() {
  const pedidos = await lerPedidos();
  const precos = await calcularPrecos(pedidos);
  const resultados = await salvarNoBanco(precos);
  resultados.forEach((msg) => console.log(msg));
}
main();
