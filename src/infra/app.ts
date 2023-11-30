import cors from "cors";
import helmet from "helmet";
import morgan from "morgan";
import express from "express";
import api from "../api/index";
import swaggerUi from "swagger-ui-express";
import swaggerDocs from "./swagger.json";

class App {
  public express: express.Application;

  constructor() {
    this.express = express();
    this.setMiddlewares();
    this.setRoutes();
    // this.catchErrors();
  }

  private setMiddlewares(): void {
    this.express.use(express.json());
    this.express.use(express.urlencoded({ extended: true }));
    this.express.use(cors());
    this.express.use(morgan("dev"));
    this.express.use(helmet());
  }

  private setRoutes(): void {
    this.express.use("/api", api);
    this.express.use("/api/docs", swaggerUi.serve, swaggerUi.setup(swaggerDocs));

  }

  private catchErrors(): void {}
}

export default new App().express;
