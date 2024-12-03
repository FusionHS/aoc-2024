package com.fusionhs.day3

import com.fusionhs.day3.answer.AnswerResolver
import io.github.oshai.kotlinlogging.KotlinLogging
import org.springframework.boot.CommandLineRunner
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.core.io.ResourceLoader

@SpringBootApplication
class Day3Application (private val answerResolver: AnswerResolver): CommandLineRunner {


	override fun run(vararg args: String?) {
		answerResolver.getAnswer(args[0]!!)
	}
}

fun main(args: Array<String>) {
	runApplication<Day3Application>(*args)
}
