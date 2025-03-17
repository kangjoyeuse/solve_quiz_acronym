object Acronym {
    fun generate(phrase: String) : String = phrase.split(", ", "-", " ", "_").filter { it.isNotEmpty() }.map { it[0] }.joinToString("").uppercase()
}