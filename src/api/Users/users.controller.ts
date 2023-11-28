import { Request, Response } from "express";

class UserController {
  async getAllUsers(req: Request, res: Response) {
    try {
      return res.status(200).json({
        message: "success",
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
