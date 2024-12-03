package com.fusionhs.day3.answer

import io.github.oshai.kotlinlogging.KotlinLogging
import org.springframework.core.io.ResourceLoader
import org.springframework.stereotype.Service
import java.nio.file.Files

private val logger = KotlinLogging.logger {}

@Service
class AnswerResolver(private val resourceLoader: ResourceLoader) {

    val patternA = Regex("""mul\((\d{1,3}),(\d{1,3})\)""")
    val patternB = Regex("""(mul)\((\d{1,3}),(\d{1,3})\)|(do)\(\)|(don't)\(\)""")

    fun getAnswer(inputFileName: String) {
        val data = loadData(inputFileName)

        runPartA(data)
        runPartB(data)

    }

    private fun runPartB(data: String) {
        val matchingPatterns = patternB.findAll(data)
        var active = true
        var sum = 0

        for (match in matchingPatterns) {

            val targetGroup = match.groupValues.filter { it.isNotEmpty() }

            when (targetGroup[1]) {
                "mul" -> if (active) {
                    sum += (match.groupValues[2].toInt() * match.groupValues[3].toInt())
                }
                "do" -> active = true
                "don't" -> active = false
                else -> throw IllegalArgumentException("Unknown match: ${match.groupValues}")
            }
        }

        logger.info { "Result B: $sum" }
    }

    private fun runPartA(data: String) {
        val matchingPatterns = patternA.findAll(data)

        val result = matchingPatterns.map { it.groupValues[1].toInt() * it.groupValues[2].toInt() }.sum()
        logger.info { "Result A: $result" }
    }

    private fun loadData(inputFileName: String): String {
        val file = resourceLoader.getResource("classpath:$inputFileName").file

        return Files.readString(file.toPath())

    }

}