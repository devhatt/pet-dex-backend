import { UserRepository } from "../repositories/User/User.repository.types";

export class UserService {
  constructor(private userRepository: UserRepository) {}

  async getAll() {
    return await this.userRepository.getAll();
  }

  async create(name: string, email: string) {
    const userExists = await this.userRepository.findByEmail(email);

    if (userExists) {
      throw new Error("Ja tem");
    }

    await this.userRepository.create(name, email);
  }
}
