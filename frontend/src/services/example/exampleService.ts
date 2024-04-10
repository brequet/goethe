export class ApiService {
  private baseUrl = "http://localhost:3000/api";

  async getHello(): Promise<string> {
    const response = await fetch(`${this.baseUrl}/hello`);
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    return response.text();
  }

  async getGoodbye(): Promise<string> {
    const response = await fetch(`${this.baseUrl}/goodbye`);
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    return response.text();
  }
}
