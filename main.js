function energiaRestante(ladrilhos) {
  const n = ladrilhos.length;
  let energia = 100;
  let currentIndex = 0;
  let visited = new Set();

  while (true) {
    if (visited.has(currentIndex)) {
      break; // Maria retornou ao ladrilho inicial, encerrando a brincadeira
    }

    visited.add(currentIndex);
    energia--; // Maria gasta 1 unidade de energia para cada pulo

    if (ladrilhos[currentIndex] === 1) {
      energia -= 2; // Se o ladrilho atual for vermelho, Maria gasta 2 unidades de energia adicionais
    }

    currentIndex = (currentIndex + ladrilhos[currentIndex]) % n;
  }

  return energia;
}

// Exemplo de uso:
const ladrilhos = [0, 1, 1, 0];
console.log(energiaRestante(ladrilhos)); // Saída: 94
