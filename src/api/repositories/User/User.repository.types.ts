import { User } from "@prisma/client";

export interface UserRepository {
  getAll(): Promise<User[]>;
  create(name: string, email: string): Promise<void>;
  findByEmail(email: string): Promise<User | null>;
}
