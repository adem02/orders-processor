# Orders Processor

---

## Run application

### Exécution

- Clonez le repo

- Pour générer le binaire:
  * La commande suivante génère un binaire nommé **orders-processor**:
    ```bash
    go build -o orders-processor main.go
    ```

  * Vous pouvez ensuite exécuter le binaire avec la commande suivante:
    ```bash
    ./order-processor <file.json>
    ```

    Avec filtre:
    ```bash
    ./order-processor -from=2024-11-01 orders.json
    ```

- Pour lancer directement avec `go run`:
  ```bash
  go run main.go <file.json>
  ```

  Avec filtre:
  ```bash
  go run main.go -from=2024-11-01 orders.json
  ```

- Si vous souhaitez traiter vos propres fichiers de commande, placez les dans le dossier **data**.

### Tests

- Pour exécuter les tests unitaires:
  ```bash
  go test ./test/* -v
  ```

---

## Mindset

### 1.

Si le programme tournait en production, on loggerait en priorité:

- Le fichier (**nomDeFichier.json**) reçu afin de s'assurer qu'on reçois bien le/les type(s) de fichiers qu'on souhaite lire et traiter, avec la taille de chaque fichier reçu
- Toutes sortes d'erreurs potentielles
- Et enfin un log de fin de fin de traitement des données et le total de commandes suspectes ainsi que celui des commandes traitées

---

### 2.

Si le fichier passait à **10Go**:

- Je passerais à une lecture de fichier en streaming. **Ainsi pas de tableau en mémoire pour les commandes**
- Pour les commandes suspectes ainsi que les revenues par marketplace, on reste dans le flux de lecture du fichier de base et on enregistre les données correspondantes dans des fichiers pendant ce flux

---

### 3.

Selon moi le cas de test prioritaire est le traitetement des données:

- total de revenues
- detection de donnée suspecte
- classement des revenues par marketplace

afin de s'assurer que le code ne retourne pas de résultats inatendus ou non prévus par le produit, ce qui pourrait conduire à des statistiques non fiables.

---
