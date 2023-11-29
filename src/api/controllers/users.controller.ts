import { Request, Response } from "express";
import { PrismaUserRepository } from "../repositories/User/User.repository";
import { UserService } from "../services/users.service";

const prismaUserRepository = new PrismaUserRepository();
const pservice = new UserService(prismaUserRepository);

export class UserController {
  async getAll(req: Request, res: Response) {
    try {
      const users = await pservice.getAll();

      return res.status(200).json(users);
    } catch (err) {}
  }
}
