import type { Config } from "@jest/types";
// Sync object
const config: Config.InitialOptions = {
  moduleNameMapper: {
    "@/(.*)": "<rootDir>/src/$1",
  },
  testPathIgnorePatterns: ["./src/test/e2e"],
  verbose: true,
  transform: {
    "^.+\\.ts?$": "ts-jest",
  },
};
export default config;
