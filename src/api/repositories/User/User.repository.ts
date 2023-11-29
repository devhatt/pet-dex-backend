import { User } from "@prisma/client";
import { prisma } from "../../../infra/db";
import { UserRepository } from "./User.repository.types";

export class PrismaUserRepository implements UserRepository {
  async create(name: string, email: string): Promise<void> {
    try {
      await prisma.user.create({
        data: {
          name,
          email,
        },
      });
    } catch (error) {
      console.error("Erro ao criar usuário:", error);
      throw error;
    }
  }

  async getAll(): Promise<User[]> {
    try {
      const users = await prisma.user.findMany();
      return users;
    } catch (error) {
      console.error("Erro ao buscar todos os usuários:", error);
      throw error;
    }
  }

  async findByEmail(email: string): Promise<User | null> {
    try {
      const user = await prisma.user.findFirst({
        where: { email },
      });

      return user;
    } catch (error) {
      console.error("Erro ao buscar usuário por e-mail:", error);
      return null;
    }
  }
}
