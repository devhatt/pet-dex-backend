import { User } from "@prisma/client";
import { UserRepository } from "../../api/repositories/User/User.repository.types";
import { randomInt } from "crypto";

export class InMemoryUserRepository implements UserRepository {
  public items: User[] = [];

  async getAll(): Promise<User[]> {
    return this.items;
  }

  async create(name: string, email: string): Promise<void> {
    this.items.push({
      id: randomInt(0, 10),
      email,
      name,
    });
  }

  async findByEmail(email: string): Promise<User | null> {
    const user = this.items.find((item) => item.email === email);

    if (!user) {
      return null;
    }

    return user;
  }
}
