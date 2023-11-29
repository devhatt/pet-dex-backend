import { Router } from "express";
import { UserController } from "../controllers/users.controller";

const users: Router = Router();
const controller = new UserController();
users.get("/", controller.getAll);

export default users;
