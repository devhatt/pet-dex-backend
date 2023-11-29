
import { UserService } from "@/api/services/users.service";
import { InMemoryUserRepository } from "./in-memory/in.memory.user.repository";

// Uma Maneira de instanciar as classes
let inMemoryUserRepo: InMemoryUserRepository;
let sut: UserService;

//  Outra maneira de instanciar. Nesse caso não precisa usar beforeEach
const makeSut = () => {
  const inMemoryUserRepo = new InMemoryUserRepository();
  const sut = new UserService(inMemoryUserRepo);

  return {
    sut,
    inMemoryUserRepo,
  };
};

// Esse é um exemplo de teste
describe("User Service", () => {
  beforeEach(() => {
    inMemoryUserRepo = new InMemoryUserRepository();
    sut = new UserService(inMemoryUserRepo);
  });

  test("Deve criar um usuário", async () => {
    await sut.create("example name", "example@email.com");

    expect(inMemoryUserRepo.items.length).toBe(1);
    expect(inMemoryUserRepo.items[0].name).toBe("example name");
  });
});
