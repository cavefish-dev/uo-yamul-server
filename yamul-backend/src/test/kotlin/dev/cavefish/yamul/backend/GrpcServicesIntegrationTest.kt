package dev.cavefish.yamul.backend

import dev.cavefish.yamul.IntegrationTest
import dev.cavefish.yamul.backend.auth.controller.BasicAuthInterceptor
import dev.cavefish.yamul.backend.character.api.CharacterServiceGrpc
import dev.cavefish.yamul.backend.character.controller.CharacterServiceController
import dev.cavefish.yamul.backend.common.api.Empty
import dev.cavefish.yamul.backend.game.controller.infra.mul.MulTileDataRepository
import dev.cavefish.yamul.backend.login.api.LoginRequest
import dev.cavefish.yamul.backend.login.api.LoginResponse.LoginResponseValue
import dev.cavefish.yamul.backend.login.api.LoginServiceGrpc
import dev.cavefish.yamul.backend.login.controller.LoginServiceController
import io.grpc.ManagedChannel
import io.grpc.Server
import io.grpc.Status
import io.grpc.StatusRuntimeException
import io.grpc.inprocess.InProcessChannelBuilder
import io.grpc.inprocess.InProcessServerBuilder
import org.assertj.core.api.Assertions.assertThat
import java.util.concurrent.TimeUnit
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.assertThrows
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.mock.mockito.MockBean

@MockBean(MulTileDataRepository::class)
class GrpcServicesIntegrationTest : IntegrationTest() {

    @Autowired private lateinit var loginServiceController: LoginServiceController
    @Autowired private lateinit var characterServiceController: CharacterServiceController
    @Autowired private lateinit var basicAuthInterceptor: BasicAuthInterceptor

    private lateinit var loginServer: Server
    private lateinit var characterServer: Server
    private lateinit var loginChannel: ManagedChannel
    private lateinit var characterChannel: ManagedChannel

    @BeforeEach
    fun setUp() {
        val loginServerName = InProcessServerBuilder.generateName()
        loginServer = InProcessServerBuilder.forName(loginServerName)
            .addService(loginServiceController)
            .directExecutor()
            .build()
            .start()
        loginChannel = InProcessChannelBuilder.forName(loginServerName).directExecutor().build()

        val characterServerName = InProcessServerBuilder.generateName()
        characterServer = InProcessServerBuilder.forName(characterServerName)
            .addService(characterServiceController)
            .intercept(basicAuthInterceptor)
            .directExecutor()
            .build()
            .start()
        characterChannel = InProcessChannelBuilder.forName(characterServerName).directExecutor().build()
    }

    @AfterEach
    fun tearDown() {
        if (::loginChannel.isInitialized) {
            loginChannel.shutdownNow()
            loginChannel.awaitTermination(5, TimeUnit.SECONDS)
        }
        if (::loginServer.isInitialized) {
            loginServer.shutdownNow()
            loginServer.awaitTermination(5, TimeUnit.SECONDS)
        }
        if (::characterChannel.isInitialized) {
            characterChannel.shutdownNow()
            characterChannel.awaitTermination(5, TimeUnit.SECONDS)
        }
        if (::characterServer.isInitialized) {
            characterServer.shutdownNow()
            characterServer.awaitTermination(5, TimeUnit.SECONDS)
        }
    }

    @Test
    fun `valid login returns valid response`() {
        val stub = LoginServiceGrpc.newBlockingStub(loginChannel)
        val response = stub.validateLogin(LoginRequest.newBuilder().setUsername("admin").build())
        assertThat(response.value).isEqualTo(LoginResponseValue.valid)
    }

    @Test
    fun `invalid login returns invalid response`() {
        val stub = LoginServiceGrpc.newBlockingStub(loginChannel)
        val response = stub.validateLogin(LoginRequest.newBuilder().setUsername("unknown").build())
        assertThat(response.value).isEqualTo(LoginResponseValue.invalid)
    }

    @Test
    fun `call without credentials returns UNAUTHENTICATED`() {
        val stub = CharacterServiceGrpc.newBlockingStub(characterChannel)
        val exception = assertThrows<StatusRuntimeException> {
            stub.getCharacterList(Empty.getDefaultInstance())
        }
        assertThat(exception.status.code).isEqualTo(Status.Code.UNAUTHENTICATED)
    }
}
