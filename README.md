# Gochat

GoChat est une application de chat simple en utilisant Go avec le framework Gin et les WebSockets pour permettre à plusieurs clients d'échanger des messages en temps réel.

![This is an image](https://github.com/gildasgatel/Gochat/blob/master/gochat.png)

## Démarrage du serveur

1. Exécutez `go run ./cmd/server` pour démarrer le serveur.
2. Ouvrez un navigateur et accédez à `http://localhost:8080` pour chaque client.

## Utilisation

1. Inscrivez votre nom dans le champ "Name".
2. Rédiger votre message.
3. Envoyer le message sur le chat avec le bouton "send".

## Fonctionnalités

- Accédez à l'application GoChat via `http://localhost:8080` une fois le serveur en route.
- Les clients peuvent écrire des messages dans le chat en même temps.
- Les messages sont échangés en temps réel à tous les clients connectés.

## Licence

Ce projet est sous licence MIT. Consultez le fichier LICENSE pour plus de détails.
