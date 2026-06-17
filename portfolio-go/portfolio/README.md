# Portfolio — Francky Relex Ndacayisaba

Site portfolio 4 pages écrit en **Go** (sans aucune dépendance externe), avec un design blanc / nuances de bleu, moderne et animé.

## Pages
| Route        | Page        | Contenu (selon le croquis) |
|--------------|-------------|-----------------------------|
| `/`          | Home        | P1 + P2 à gauche, photo à droite, **bande défilante** en bas |
| `/about`     | About me    | Paragraphes P3 → P7 en zigzag |
| `/portfolio` | Portfolio   | Title 1 (stages), Title 2.1 / 2.2 (projets), Title 3 (certifications) |
| `/contact`   | Contact me  | P8 + cartes Téléphone, Email, LinkedIn, GitHub, Instagram |

## Lancer le site
```bash
go run .
# puis ouvrir http://localhost:8080
```
(Nécessite Go 1.22+. Le port se change avec la variable `PORT`.)

## Personnaliser
- **Ta photo** : ajoute `static/img/photo.jpg`, puis dans `templates/home.html`
  remplace le bloc `.photo-placeholder` par :
  ```html
  <img src="/static/img/photo.jpg" alt="Francky Relex Ndacayisaba">
  ```
- **Textes** : modifie directement les fichiers dans `templates/`
  (P1/P2 dans `home.html`, P3–P7 dans `about.html`, P8 dans `contact.html`).
- **Liens de contact** : mets tes vrais numéro, email et profils dans `templates/contact.html`.
- **Couleurs / polices** : tout est centralisé dans les variables `:root`
  en haut de `static/css/style.css`.
- **Compétences de la bande défilante** : liste dans `templates/home.html`
  (pense à modifier les **deux** groupes identiques, c'est ce qui rend le défilement infini).

## Structure
```
portfolio/
├── main.go              # serveur HTTP + routage
├── go.mod
├── templates/           # base.html + 1 fichier par page
└── static/
    ├── css/style.css    # design complet (responsive + animations)
    └── js/main.js       # révélations au scroll + menu mobile
```
