package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type No struct {
	hashNo []byte

	esqueda *No
	direita *No
}

func encontrar_hash(passo string) []byte {
	h := sha256.Sum256([]byte(passo))
	return h[:]
}

func criar_folhas(hashes [][]byte) []*No {
	folhas := make([]*No, 0)

	for _, hash := range hashes {
		folhas = append(folhas, &No{
			hashNo:  hash,
			esqueda: nil,
			direita: nil,
		})
	}

	return folhas
}

func criar_intermediarios(nos []*No) []*No {
	intermediarios := make([]*No, 0)

	// aqui iremos ir de 2 em 2 folhas até o final
	for i := 0; i < len(nos); i += 2 {
		folhaAtual := nos[i]

		// a folhaIrma faz referência a folha que esta
		// ao lado, no caso de uma lista com numero impar
		// uma folha sempre ficará sozinha
		// exemplo: folhas A - B - C, vamos combinar A com B
		// e C ficará sozinha, então vamos combinas C com C
		var folhaIrma *No
		if i+1 >= len(nos) {
			folhaIrma = nos[i]
		} else {
			folhaIrma = nos[i+1]
		}

		// combino o hash das 2 e gero um novo hash
		combinaHash := bytes.Join(
			[][]byte{folhaAtual.hashNo, folhaIrma.hashNo}, []byte{})
		novoHash := sha256.Sum256(combinaHash)

		noIntermediario := &No{
			hashNo:  novoHash[:],
			esqueda: folhaAtual,
			direita: folhaIrma,
		}

		intermediarios = append(intermediarios, noIntermediario)
	}

	return intermediarios
}

func combinar_intermediarios(intermediarios []*No) *No {
	for {
		quantidade_intermediarios := len(intermediarios)

		if quantidade_intermediarios == 1 {
			return intermediarios[0]
		}

		intermediarios = criar_intermediarios(intermediarios)
	}
}

func criar_raiz(hashes [][]byte) *No {
	folhas := criar_folhas(hashes)
	intermediarios := criar_intermediarios(folhas)

	raiz := combinar_intermediarios(intermediarios)
	return raiz
}

func main() {
	fazer_bolo_contrato := []string{
		"pegar ovos, açucar, oléo e cenoura",
		"bater tudo no liquidificador",
		"misturar com trigo e fermento",
		"bater a massa por 3m",
		"untar a forma",
		"jogar a massa na forma",
		"preaquecer o forno à 180 graus por 10m",
		"assar a massa que está na forma por 45m",
	}

	fazer_bolo_hashes := make([][]byte, 0)

	for _, passo := range fazer_bolo_contrato {
		fazer_bolo_hashes = append(fazer_bolo_hashes, encontrar_hash(passo))
	}

	raiz := criar_raiz(fazer_bolo_hashes)

	fmt.Printf("%x\n", raiz.hashNo)
}
