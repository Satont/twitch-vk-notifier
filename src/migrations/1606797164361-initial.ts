import {MigrationInterface, QueryRunner} from "typeorm";

export class initial1606797164361 implements MigrationInterface {
    name = 'initial1606797164361'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`
            CREATE TABLE "chats" (
                "id" character varying NOT NULL,
                "followGameChange" boolean NOT NULL,
                CONSTRAINT "PK_0117647b3c4a4e5ff198aeb6206" PRIMARY KEY ("id")
            )
        `);
        await queryRunner.query(`
            CREATE TABLE "follows" (
                "id" SERIAL NOT NULL,
                "chatId" character varying,
                "channelId" character varying,
                CONSTRAINT "PK_8988f607744e16ff79da3b8a627" PRIMARY KEY ("id")
            )
        `);
        await queryRunner.query(`
            CREATE TABLE "channels" (
                "id" character varying NOT NULL,
                "username" character varying NOT NULL,
                "online" boolean NOT NULL,
                "game" character varying NOT NULL,
                "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
                "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
                CONSTRAINT "PK_bc603823f3f741359c2339389f9" PRIMARY KEY ("id")
            )
        `);
        await queryRunner.query(`
            ALTER TABLE "follows"
            ADD CONSTRAINT "FK_dc1c79f16fa7bf1ef500a0305fb" FOREIGN KEY ("chatId") REFERENCES "chats"("id") ON DELETE NO ACTION ON UPDATE NO ACTION
        `);
        await queryRunner.query(`
            ALTER TABLE "follows"
            ADD CONSTRAINT "FK_af75a484b4316e01b2248a7d244" FOREIGN KEY ("channelId") REFERENCES "channels"("id") ON DELETE NO ACTION ON UPDATE NO ACTION
        `);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`
            ALTER TABLE "follows" DROP CONSTRAINT "FK_af75a484b4316e01b2248a7d244"
        `);
        await queryRunner.query(`
            ALTER TABLE "follows" DROP CONSTRAINT "FK_dc1c79f16fa7bf1ef500a0305fb"
        `);
        await queryRunner.query(`
            DROP TABLE "channels"
        `);
        await queryRunner.query(`
            DROP TABLE "follows"
        `);
        await queryRunner.query(`
            DROP TABLE "chats"
        `);
    }

}
