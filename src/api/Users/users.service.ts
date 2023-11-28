import { Request, Response } from "express";

class UserService {
  async getAll(req: Request) {
    return "sucess";
  }

  // Outros métodos do UserService podem ser adicionados aqui
}

const userService = new UserService();

export default userService;
