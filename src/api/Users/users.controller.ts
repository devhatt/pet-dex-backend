import { Request, Response } from "express";
import userService from "./users.service";

class UserController {
  async getAllUsers(req: Request, res: Response) {
    try {
      const result = await userService.getAll(req);
      return res.status(200).json({
        message: result,
      });
    } catch (error) {
      console.error(error);
      res.status(500).send({
        message: error,
      });
    }
  }

  // Outros métodos do UserController podem ser adicionados aqui
}

const userController = new UserController();

export default userController;
