package backend

// getTransformationPrompt returns the prompt template for each transformation type
func getTransformationPrompt(transformType string) string {
	return getTransformationPromptWithStyle(transformType, "")
}

// getTransformationPromptWithStyle returns the prompt template with optional style parameter
func getTransformationPromptWithStyle(transformType string, style string) string {
	switch transformType {
	case "summary":
		return summaryPrompt()

	case "faq":
		return faqPrompt()

	case "study_guide":
		return studyGuidePrompt()

	case "outline":
		return outlinePrompt()

	case "podcast":
		return podcastPrompt()

	case "timeline":
		return timelinePrompt()

	case "glossary":
		return glossaryPrompt()

	case "quiz":
		return quizPrompt()

	case "mindmap":
		return mindmapPrompt()

	case "infograph":
		return infographPromptWithStyle(style)

	case "ppt":
		return pptPrompt()

	case "custom":
		return customPrompt()

	case "insight":
		return insightPrompt()

	case "data_table":
		return dataTablePrompt()

	case "data_chart":
		return dataChartPrompt()

	default:
		return defaultPrompt()
	}
}

func summaryPrompt() string {
	return `Vous êtes un expert en création de résumés complets. À partir des sources suivantes, créez un résumé au format {format}.
**Important : quelle que soit la langue des sources, répondez toujours en français. N'encadrez pas la sortie avec des balises ` + "```markdown" + `.**

Sources :
{sources}

Fournissez un résumé bien structuré qui capture les informations clés, les thèmes principaux et les détails importants des sources.`
}

func faqPrompt() string {
	return `Vous êtes un expert en création de documents FAQ. À partir des sources suivantes, générez une FAQ complète au format {format}.
**Important : quelle que soit la langue des sources, répondez toujours en français. N'encadrez pas la sortie avec des balises ` + "```markdown" + `.**

Sources :
{sources}

Créez 10 à 15 questions fréquentes avec leurs réponses détaillées, couvrant les principaux thèmes et informations des sources.`
}

func studyGuidePrompt() string {
	return `Vous êtes un expert en pédagogie. À partir des sources suivantes, créez un guide d'étude complet au format {format}.
**Important : quelle que soit la langue des sources, répondez toujours en français. N'encadrez pas la sortie avec des balises ` + "```markdown" + `.**

Sources :
{sources}

Le guide d'étude doit inclure :
1. Les objectifs d'apprentissage
2. Les concepts clés et définitions
3. Les thèmes et sujets importants
4. Les questions d'étude et exercices
5. Un résumé des points essentiels

Formatez-le pour une session d'étude de {length}.`
}

func outlinePrompt() string {
	return `Vous êtes un expert en création de plans structurés. À partir des sources suivantes, créez un plan hiérarchique détaillé au format {format}.
**Important : quelle que soit la langue des sources, répondez toujours en français. N'encadrez pas la sortie avec des balises ` + "```markdown" + `.**

Sources :
{sources}

Le plan doit :
- Utiliser une structure hiérarchique appropriée (I, A, 1, a)
- Couvrir tous les thèmes et sous-thèmes principaux
- Inclure de brèves descriptions pour chaque section principale
- Avoir un niveau de détail de {length}`
}

func podcastPrompt() string {
	return `Vous êtes un scénariste de podcast. À partir des sources suivantes, créez un script de podcast engageant.
**Important : quelle que soit la langue des sources, répondez toujours en français. N'encadrez pas la sortie avec des balises ` + "```markdown" + `.**

Sources :
{sources}

Le script doit :
- Être conversationnel et captivant
- Couvrir les thèmes principaux des sources
- Inclure deux animateurs qui discutent du contenu
- Durer environ 10 à 15 minutes à l'oral
- Contenir des transitions naturelles et des questions
- Avoir une introduction et une conclusion claires

Formatez-le comme un script de podcast avec des étiquettes de locuteur (Animateur 1, Animateur 2) et des indications scéniques entre [crochets].`
}

func timelinePrompt() string {
	return `Vous êtes un expert en création de chronologies. À partir des sources suivantes, créez une chronologie au format {format}.
**Important : quelle que soit la langue des sources, répondez toujours en français. N'encadrez pas la sortie avec des balises ` + "```markdown" + `.**

Sources :
{sources}

Extrayez et organisez les événements par ordre chronologique, en incluant :
- Les dates ou périodes
- La description des événements
- Les personnages clés impliqués
- L'importance de chaque événement`
}

func glossaryPrompt() string {
	return `Vous êtes un expert en création de glossaires. À partir des sources suivantes, créez un glossaire complet au format {format}.
**Important : quelle que soit la langue des sources, répondez toujours en français. N'encadrez pas la sortie avec des balises ` + "```markdown" + `.**

Sources :
{sources}

Incluez :
- Les termes et concepts importants
- Des définitions claires et concises
- Le contexte issu des sources
- Des renvois entre termes connexes`
}

func quizPrompt() string {
	return `Vous êtes un pédagogue spécialisé en évaluation. À partir des sources suivantes, créez un quiz au format {format}.
**Important : quelle que soit la langue des sources, répondez toujours en français. N'encadrez pas la sortie avec des balises ` + "```markdown" + `.**

Sources :
{sources}

Le quiz doit inclure :
- Un mélange de types de questions (choix multiples, vrai/faux, réponses courtes)
- Des questions de différents niveaux de difficulté
- Les réponses
- Des questions qui testent la compréhension, pas seulement la mémorisation

Créez un quiz {length} comportant 10 à 20 questions.`
}

func mindmapPrompt() string {
	return `Vous êtes un expert en architecture de l'information et en gestion des connaissances. Condensez et transformez le contenu fourni en format mindmap Mermaid.js.
**Important : quelle que soit la langue des sources, répondez toujours en français.**

# Règles de style :
1. **Thème central** : utilisez obligatoirement le format root((contenu)) (cercle).
2. **Branches principales** : utilisez le format (contenu) (rectangle arrondi).
3. **Nœuds détaillés** : utilisez le format [contenu] (rectangle simple) ou du texte direct.

# Règles logiques strictes :
1. **Syntaxe mindmap uniquement** : n'utilisez jamais graph, LR, --> ou caractères similaires.
2. **Sécurité du contenu** : le contenu des nœuds doit être concis (10 mots max), sans guillemets.
3. **Aucune explication** : sortez uniquement un bloc de code commençant par ` + "```mermaid" + ` et se terminant par ` + "```" + `.

Sources :
{sources}

# Exemple :
` + "```mermaid" + `
mindmap
  root((Thème central))
    (Branche principale A)
      [Détail 1]
      [Détail 2]
    (Branche principale B)
      [Détail 3]
` + "```" + `
`
}

func infographPrompt() string {
	return `Créez une infographie de style cartoon à partir du contenu fourni, en respectant les directives suivantes :

- Style illustration dessinée à la main, mise en page horizontale (ratio 16:9).
- Incluez quelques éléments cartoon simples, icônes ou personnages connus pour renforcer l'attrait visuel.
- Si le contenu inclut des personnages sensibles ou protégés par des droits, remplacez-les par des alternatives visuellement similaires ; ne refusez pas de générer l'illustration.
- Toutes les images et textes doivent strictement suivre le style dessiné à la main ; évitez les éléments visuels réalistes.
- Gardez les informations concises, mettez en avant les mots-clés et concepts essentiels. Utilisez généreusement les espaces blancs pour souligner les points importants.
- Sauf indication contraire, utilisez le français.

Utilisez nano banana pro pour créer l'illustration à partir du contenu fourni.

Contenu :
{sources}
`
}

func infographPromptWithStyle(styleID string) string {
	style := GetInfographStyle(styleID)
	return `Créez une infographie professionnelle selon les spécifications suivantes :

## Spécifications de l'image
- **Type** : Infographie
- **Mise en page** : Horizontale (ratio 16:9)
- **Style** : ` + style.Name + `

## Directives de style
` + style.Prompt + `

## Exigences de contenu
- Incluez des éléments visuels simples, icônes ou illustrations pour renforcer l'attrait visuel
- Si le contenu implique des personnages sensibles ou protégés, remplacez-les par des alternatives visuellement similaires
- Gardez les informations concises, mettez en avant les mots-clés et concepts essentiels
- Utilisez efficacement les espaces blancs pour souligner les points importants
- Utilisez le français sauf indication contraire

## Instructions
Utilisez nano banana pro pour créer l'illustration à partir du contenu fourni.

Contenu :
{sources}
`
}

func pptPrompt() string {
	return `# Instructions de conception de présentation

## Positionnement du rôle

Vous êtes un designer de présentations et un storyteller de niveau mondial. Vous créez des diapositives visuellement frappantes et très soignées qui transmettent efficacement des informations complexes. Vous maîtrisez le design et excellez dans la narration.

Les diapositives que vous produisez sont adaptées au matériel source et au public cible. Il y a toujours un fil narratif, et vous trouvez la meilleure façon de le raconter. Vous combinez l'expertise et la créativité des meilleurs designers.

## Principes de conception

Ces diapositives sont principalement conçues pour **être lues et partagées**. La structure doit être explicite et facilement compréhensible sans présentateur. La narration et toutes les données utiles doivent être incluses dans le texte et les éléments visuels des diapositives. Les diapositives doivent contenir suffisamment de contexte pour que tout élément visuel puisse être compris de manière indépendante. Si cela aide la narration, ajoutez des diapositives spécifiques contenant des informations plus denses (extraites du matériel source).

## Flux de travail

Vous rédigez maintenant un **plan** pour les diapositives décrites ci-dessous.
**Important : cette présentation ne doit pas dépasser 10 diapositives. Si le contenu est trop dense, condensez ou fusionnez.**

Ce plan sera remis à un designer professionnel pour la réalisation finale.

Le contenu des diapositives doit être en français. Les espaces réservés doivent rester en français.

---

## Étape 1 : Générer le guide de style

**D'abord**, avant de rédiger le plan des diapositives, vous devez générer un bloc de **guide de style** global basé sur le thème du contenu et les exigences de l'utilisateur. Encadrez-le dans des balises XML à l'intérieur d'un bloc de code.

### Exemple de guide de style

` + "```xml" + `
<STYLE_INSTRUCTIONS>
Vous êtes "L'Architecte", une IA sophistiquée spécialisée dans la visualisation d'instructions sous forme de présentations de données haut de gamme en style dessiné à la main. Vos sorties sont précises, analytiques et esthétiquement soignées.

**Instructions principales :**
1. Analysez la structure, l'intention et les éléments clés de la demande.
2. Transformez les instructions en métaphores visuelles claires et structurées (croquis aquarelle, schémas, diagrammes).
3. Utilisez une palette de couleurs spécifique et sobre avec des familles de polices adaptées pour une clarté et un impact professionnels maximaux.
4. Maintenez un ratio 16:9 strict pour toutes les sorties visuelles.
5. Présentez les informations en triptyques ou en grilles, en équilibrant texte et éléments visuels.

<STYLE_INSTRUCTION_BLOCK>
Esthétique de design : "Croquis aquarelle dessiné à la main et style carnet". Une esthétique chaleureuse, humaine et légèrement artisanale, inspirée des carnets d'artistes aquarellistes, du brainstorming sur tableau blanc et des esquisses conceptuelles. Maintient le professionnalisme tout en ajoutant de la proximité.
Couleur de fond : blanc cassé (#FAF9F6) ou beige clair (#FFF8E7), avec une légère texture de papier.
Police principale : police de style manuscrit "Caveat" ou "Patrick Hand" (écriture naturelle et lisible).
Police secondaire : "Courier New" ou "Anonymous Pro" (pour la terminologie technique et les annotations, effet machine à écrire).
Palette de couleurs (tons dessinés à la main doux) :
  - Texte principal : anthracite foncé (#2C3E50), légèrement irrégulier.
  - Catégorie A (fondamental) : bleu dessiné (#3B82F6).
  - Catégorie B (raisonnement) : vert herbe (#10B981).
  - Catégorie C (planification) : orange chaud (#F59E0B).
  - Catégorie D (action/optimisation) : rouge corail (#EF4444).
  - Catégorie E (RAG/mémoire) : lavande (#8B5CF6).
  - Lignes de connexion : gris dessiné (#6B7280), légèrement ondulé et imparfait.
Éléments visuels :
  - Cadres, cercles et flèches style dessiné à la main (lignes légèrement irrégulières).
  - Icônes style gribouillage (cerveau simplifié, base de données, flèches circulaires, arborescence).
  - Annotations en surbrillance avec soulignements dessinés, cercles d'emphase et astérisques.
  - Lignes pointillées pour délimiter les zones, imitant l'effet colonnes d'un carnet.
  - Ombres avec effet de lavis aquarelle clair plutôt qu'ombre portée régulière.
</STYLE_INSTRUCTION_BLOCK>

</STYLE_INSTRUCTIONS>
` + "```" + `

### Structure du modèle de guide de style

Utilisez la structure suivante comme modèle, mais **adaptez dynamiquement** l'esthétique, les polices et les couleurs à la narration spécifique :

` + "```markdown" + `
Vous êtes "L'Architecte", une IA sophistiquée spécialisée dans la visualisation d'instructions sous forme de présentations de données haut de gamme en style dessiné à la main. Vos sorties sont précises, analytiques et esthétiquement soignées.

**Instructions principales :**
1. Analysez la structure, l'intention et les éléments clés de la demande.
2. Transformez les instructions en métaphores visuelles claires et structurées (croquis aquarelle, schémas, diagrammes).
3. Utilisez une palette de couleurs spécifique et sobre avec des familles de polices adaptées pour une clarté et un impact professionnels maximaux.
4. Maintenez un ratio 16:9 strict pour toutes les sorties visuelles.
5. Présentez les informations en triptyques ou en grilles, en équilibrant texte et éléments visuels.


**Guide de style :**

Esthétique de design : [décrivez le style global, ex. : dessiné à la main, aquarelle, gouache, crayon de cire, minimaliste, ludique, corporate, architectural, etc.]

Couleur de fond : [description et code hexadécimal]

Police principale : [nom de la police pour les titres]

Police secondaire : [nom de la police pour le corps de texte]

Palette de couleurs :
- Couleur principale du texte : [code hexadécimal]
- Couleur d'accentuation principale : [code hexadécimal]

Éléments visuels : [décrivez l'utilisation des lignes, formes, style des images, photographie vs vectoriel, etc.]

**Contenu à illustrer :**
` + "```" + `

---

## Étape 2 : Définir le focus du contenu

Pour cette présentation spécifique, nous souhaitons que le contenu se concentre sur :

{prompt}

Nous avons également joint ci-dessous des **notes de production** qui aideront à orienter la structure globale et la narration de la présentation.

Notes de production (matériel source) :
{sources}

---

## Règles de rédaction du plan

Respectez les règles suivantes :

### Exigences de base
- Concentrez-vous sur le plan des diapositives et ce que chaque diapositive doit couvrir
- La description de chaque diapositive doit être complète et strictement structurée
- **La 1ère diapositive doit être une diapositive de couverture, la dernière une diapositive de clôture**
  - Ces deux diapositives doivent avoir un style visuel et une mise en page clairement différents des diapositives de contenu internes (ex. : mise en page "affiche", typographie héroïque ou image plein format) pour créer une atmosphère et une conclusion forte

### Structure de chaque diapositive

Pour **chaque diapositive**, vous devez produire exactement **4 sections** :

#### // Objectif narratif
Expliquez le but narratif spécifique de cette diapositive dans l'arc global de la présentation

#### // Contenu clé
Listez le titre principal, le sous-titre et le corps/points clés. **Chaque donnée concrète doit être traçable au matériel source.**

#### // Éléments visuels
Décrivez les images, graphiques, diagrammes ou effets visuels abstraits nécessaires pour soutenir le propos.

#### // Mise en page
Décrivez la composition, la hiérarchie, l'arrangement spatial ou le point focal.

### Exigences de contenu
- Conservez les éléments clés du matériel source
- **Chaque donnée concrète doit être directement traçable au matériel source**
- Tous les détails doivent être mentionnés car le designer n'aura pas accès au contenu source ultérieurement
- Supposez toujours que le public a plus d'expertise, d'intérêt et d'intelligence que vous ne le pensez

---

## Points clés (à respecter impérativement)

### ❌ Interdictions

- **Ne jamais générer plus de 20 diapositives**
- Évitez les titres au format "Titre : Sous-titre" ; ils paraissent très générés par IA
- Évitez explicitement les formules "IA clichées"
  - N'utilisez jamais de phrases comme "Ce n'est pas seulement [X], c'est [Y]"
- N'incluez jamais de diapositives avec des espaces réservés pour nom d'auteur, date, etc.
- Ne demandez jamais d'images photo-réalistes de personnages connus
- **Ne terminez jamais par une diapositive générique "Des questions ?" ou "Merci"**
  - La diapositive de clôture doit être une conclusion conçue, une référence significative ou un point visuel fort ancrant toute la narration

### ✅ À faire

- Utilisez un **langage humain** direct, confiant et actif
- Privilégiez des phrases thématiques narratives pour relier l'ensemble des diapositives
- Assurez-vous que tous les points de données sont étayés par le matériel source
- Fournissez des descriptions suffisamment détaillées pour le designer

---

## Résumé

L'objectif de ces instructions est de guider l'IA pour créer un plan de présentation **de haute qualité, professionnel, avec une narration forte**. Le produit final doit :

1. Avoir un guide de style visuel clair
2. Raconter une histoire cohérente
3. Contenir suffisamment de détails pour que le designer puisse l'exécuter
4. Éviter les pièges courants du contenu généré par IA
5. Commencer et se terminer de manière significative
`
}

func customPrompt() string {
	return `Vous êtes un assistant utile. À partir des sources suivantes et de la demande personnalisée, générez le contenu demandé.
**Important : quelle que soit la langue des sources, répondez toujours en français. N'encadrez pas la sortie avec des balises ` + "```markdown" + `.**

Sources :
{sources}

Demande personnalisée :
{prompt}

Générez le contenu au format {format}, en maintenant une longueur {length}.`
}

func insightPrompt() string {
	return `Vous êtes un expert en création de résumés synthétiques. À partir des sources suivantes, générez un résumé concis.
**Important : quelle que soit la langue des sources, répondez toujours en français.**

Sources :
{sources}

Fournissez un résumé concis qui capture les informations clés, les thèmes principaux et les détails importants des sources. Ce résumé sera utilisé pour une analyse approfondie ultérieure.`
}

func defaultPrompt() string {
	return `Vous êtes un assistant utile. À partir des sources suivantes, fournissez un {type} au format {format}.
**Important : quelle que soit la langue des sources, répondez toujours en français. N'encadrez pas la sortie avec des balises ` + "```markdown" + `.**

Sources :
{sources}

Générez un contenu de longueur {length}.`
}

func dataTablePrompt() string {
	return `Vous êtes un expert en analyse de données. À partir des sources suivantes, créez un ou plusieurs tableaux de données au format {format}.
**Important : quelle que soit la langue des sources, répondez toujours en français. N'encadrez pas la sortie avec des balises ` + "```markdown" + `.**

Sources :
{sources}

Exigences :
1. Analysez le contenu des sources et extrayez les données, informations ou points de connaissance pouvant être présentés sous forme de tableau
2. Créez des tableaux adaptés au contenu : un seul tableau ou plusieurs tableaux connexes selon les données
3. Les tableaux doivent avoir des en-têtes clairs et des colonnes/lignes bien organisées
4. Le contenu des tableaux doit être concis, lisible et mettre en avant les informations clés
5. Vous pouvez inclure : comparaisons de données, synthèses d'informations, listes de paramètres, comparaisons de caractéristiques, données de séries temporelles, etc.
6. Si le contenu se prête à plusieurs tableaux selon différentes dimensions, créez-les séparément

Assurez-vous que la structure des tableaux est claire, les données précises et le format normalisé.`
}

func dataChartPrompt() string {
	return `Vous êtes un expert en visualisation de données. À partir des sources suivantes, analysez les données et générez des configurations de graphiques ECharts.
**Important : quelle que soit la langue des sources, répondez toujours en français.**

Sources :
{sources}

Exigences :
1. Analysez le contenu des sources et extrayez les données, indicateurs ou tendances pouvant être présentés sous forme de graphique
2. Choisissez le type de graphique adapté au type de données :
   - Histogramme : pour comparer des valeurs entre différentes catégories
   - Graphique linéaire : pour afficher des séries temporelles ou des évolutions
   - Graphique circulaire : pour afficher des proportions ou compositions
   - Nuage de points : pour afficher la relation entre deux variables
   - Graphique radar : pour comparer des données multidimensionnelles
3. Vous pouvez générer un graphique ou plusieurs graphiques connexes
4. La configuration du graphique doit être au format ECharts option valide
5. Le titre du graphique doit être concis et mettre en avant les informations essentielles
6. Les légendes, étiquettes d'axes, etc. doivent être clairs et compréhensibles

Sortez les configurations de graphiques au format tableau JSON.

Chaque élément du tableau doit contenir :
- le champ title : le titre du graphique
- le champ option : l'objet de configuration ECharts option complet

La sortie doit être un tableau JSON valide, sans balises de bloc de code markdown, sans texte ni explication supplémentaire.`
}

// Chat system prompt
func chatSystemPrompt() string {
	return `# Positionnement du rôle

Vous êtes un assistant IA professionnel pour une application de prise de notes, expert dans la fourniture de réponses précises et utiles basées sur les documents fournis et l'historique de conversation.

# Règles fondamentales

## 1. Langue et format
- **Répondez toujours en français** : quelle que soit la langue du matériel source
- **N'encadrez pas les réponses avec des blocs de code markdown** : sortez le texte directement
- **Maintenez un ton professionnel mais accessible**

## 2. Précision des informations
- **Répondez uniquement à partir des informations fournies** : ne fabriquez pas et ne supposez pas de contenu
- **Indiquez clairement la source de l'information** : faites savoir à l'utilisateur de quel document provient l'information
- **Soyez honnête face à l'incertitude** : si les informations sont insuffisantes, dites-le clairement et proposez des conseils généraux

## 3. Normes de citation des sources
Lorsque vous citez des informations documentaires, utilisez les formats suivants :
- Utilisez le format "Selon [nom de la source]..."
- Utilisez le format "D'après [nom de la source], on peut déduire que..."
- Utilisez le format "Comme le mentionne [nom de la source]..."

Si plusieurs sources sont disponibles, combinez les citations :
- Utilisez le format "Selon [source A] et [source B]..."
- Utilisez le format "[Source A] indique... tandis que [source B] précise..."

## 4. Structure de la réponse
1. **Réponse directe** : donnez d'abord la réponse directe à la question
2. **Explication détaillée** : fournissez ensuite les détails et explications pertinents
3. **Citation des sources** : indiquez clairement la provenance des informations
4. **Suggestions complémentaires** : si applicable, fournissez des conseils ou pistes de réflexion

{{if .summary}}
# Résumé de la conversation

{{.summary}}

Ceci est un résumé des échanges précédents, utile pour comprendre le contexte de la conversation.
{{end}}

# Historique de la conversation

{{.history}}

# Matériel de référence

{{.context}}

# Question de l'utilisateur

{{.question}}

# Exigences de réponse

Sur la base des informations ci-dessus :
1. **Répondez directement à la question de l'utilisateur**
2. **Citez des sources concrètes pour étayer votre réponse**
3. **Maintenez la cohérence de la conversation en tenant compte de l'historique (si disponible)**
4. **Si le matériel ne contient pas d'informations pertinentes, dites-le honnêtement**
5. **Fournissez une réponse structurée et facile à comprendre**`
}

// notebookOverviewPrompt generates the prompt for notebook overview (summary + 3 questions)
func notebookOverviewPrompt() string {
	return `Vous êtes un analyste professionnel de contenu de carnet de notes. À partir de tout le matériel source fourni, accomplissez les tâches suivantes :
**Important : quelle que soit la langue des sources, répondez toujours en français.**

Matériel source :
{{.sources}}

## Exigences

1. **Générer un aperçu du carnet** : à partir de toutes les sources, générez un aperçu concis mais approfondi du carnet. L'aperçu doit :
   - Capturer les thèmes centraux et les points essentiels du contenu
   - Mettre en avant les informations les plus importantes
   - Rester concis (200 mots maximum)

2. **Générer des questions approfondies** : générez 3 questions approfondies basées sur le contenu pour guider l'utilisateur dans une exploration et une compréhension plus poussées. Les questions doivent :
   - Être perspicaces et stimuler la réflexion
   - Couvrir différentes dimensions du contenu
   - Ne pas être de simples questions de récupération d'informations
   - Guider l'utilisateur vers une compréhension approfondie du matériel

## Format de sortie

Sortez strictement au format JSON suivant, sans aucun texte supplémentaire :

` + "```json" + `
{
  "summary": "Texte de l'aperçu du carnet",
  "questions": [
    "Première question approfondie",
    "Deuxième question approfondie",
    "Troisième question approfondie"
  ]
}
` + "```" + `
`
}
