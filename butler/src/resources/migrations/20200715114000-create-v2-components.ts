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

import { MigrationInterface, QueryRunner } from 'typeorm'

export class CreateV2Components20200715114000 implements MigrationInterface {

  public async up(queryRunner: QueryRunner) : Promise<void> {
    await queryRunner.query(`
      CREATE TABLE "public"."v2components" (
        "id" uuid DEFAULT uuid_generate_v4 () NOT NULL,
        "deployment_id" uuid DEFAULT uuid_generate_v4 () NOT NULL,
        "component_id" uuid NOT NULL,
        "name" Character Varying NOT NULL,
        "image_url" Character Varying NOT NULL,
        "image_tag" Character Varying NOT NULL,
        "helm_url" Character Varying NOT NULL,
        "running" Boolean DEFAULT false NOT NULL,
        "merged" Boolean DEFAULT false NOT NULL,
        "created_at" timestamp without time zone DEFAULT now() NOT NULL,
        PRIMARY KEY ( "id" ),
        CONSTRAINT "fk_v2deployments" FOREIGN KEY ( "deployment_id" ) REFERENCES "public"."v2deployments" ( "id" )
      );

      CREATE INDEX "index_components_id" ON "public"."v2components" USING btree( "deployment_id" );
      CREATE INDEX "index_name" ON "public"."v2components" USING btree( "name" );
      CREATE UNIQUE INDEX "only_one_module_running" ON "public"."v2components" (running, name) WHERE (running);
      `)
  }

  public async down(queryRunner: QueryRunner) : Promise<void> {
    await queryRunner.dropTable('v2components')
  }
}
