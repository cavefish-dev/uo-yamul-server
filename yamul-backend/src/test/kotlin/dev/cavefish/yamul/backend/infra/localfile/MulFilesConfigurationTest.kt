package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.IntegrationTest
import dev.cavefish.yamul.backend.Constants
import dev.cavefish.yamul.backend.infra.localfile.MultimaFileRepository.IndexedFileProperties
import dev.cavefish.yamul.backend.infra.localfile.MultimaFileRepository.PlainMulFileProperties
import dev.cavefish.yamul.backend.infra.localfile.MultimaFileRepository.UopFileProperties
import org.assertj.core.api.Assertions.assertThat
import java.io.File
import kotlin.test.Test

class MulFilesConfigurationTest : IntegrationTest() {

    @Test
    fun mulFilesFolderExistsAndContainsExpectedFiles() {
        val path = System.getenv(Constants.MULTIMA_PATH)
        assertThat(path)
            .describedAs("${Constants.MULTIMA_PATH} environment variable must be set")
            .isNotBlank()

        val dir = File(path!!)
        assertThat(dir)
            .describedAs("MUL files directory '$path'")
            .exists()
            .isDirectory()

        buildExpectedFileList().forEach { filename ->
            softly.assertThat(dir.resolve(filename))
                .describedAs("Expected MUL file: $filename")
                .exists()
                .isFile()
        }
    }

    private fun buildExpectedFileList(): List<String> {
        val files = mutableListOf("tiledata.mul")
        MulMapHelper.mapProperties.forEach { map ->
            files.addAll(filesFrom(map.mapFile))
            files.addAll(filesFrom(map.staticsFile))
        }
        return files
    }

    private fun filesFrom(props: MultimaFileRepository.MulFileProperties): List<String> =
        when (props) {
            is UopFileProperties -> props.filenames
            is PlainMulFileProperties -> listOf(props.filename)
            is IndexedFileProperties -> listOf(props.baseFilename, props.idxFilename)
        }
}
