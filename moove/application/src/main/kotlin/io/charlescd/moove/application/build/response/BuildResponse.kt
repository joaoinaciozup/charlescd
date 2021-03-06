/*
 * Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.charlescd.moove.application.build.response

import com.fasterxml.jackson.annotation.JsonFormat
import io.charlescd.moove.application.user.response.SimpleUserResponse
import io.charlescd.moove.domain.Build
import java.time.LocalDateTime

data class BuildResponse(
    val id: String,
    val author: SimpleUserResponse,
    @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd HH:mm:ss")
    val createdAt: LocalDateTime,
    val features: List<FeatureResponse>,
    val tag: String,
    val status: String,
    val deployments: List<SimpleDeploymentResponse>
) {

    companion object {
        fun from(build: Build): BuildResponse {
            return BuildResponse(
                id = build.id,
                author = SimpleUserResponse.from(build.author),
                createdAt = build.createdAt,
                features = build.features.map { FeatureResponse.from(it) },
                tag = build.tag,
                status = build.status.name,
                deployments = build.deployments.map { SimpleDeploymentResponse.from(it) }
            )
        }
    }
}
