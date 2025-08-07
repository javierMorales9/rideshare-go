
const url = "http://localhost:8080"
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQ2NzgxODcsInVzZXJfaWQiOjF9.n9LSSWw4yzubadAuCq3s3V4Q9rGjFev5XNN1Zw3wRaQ";

async function register(payload) {
  const result = await fetch(`${url}/auth/register`, {
    method: "POST",
    body: JSON.stringify(payload),
  });

  return await res.text();
}

async function login(payload) {
  const res = await fetch(`${url}/auth/login`, {
    method: "POST",
    body: JSON.stringify(payload),
  });

  return await res.text();
}

async function me() {
  const res = await fetch(`${url}/api/me`, {
    headers: {
      Authorization: `Bearer ${token}`,
    }
  });
  return await res.json();
}

const user = {
  first_name: "Pedrito",
  last_name: "Lanzadera",
  email: "pedrolanza@gmail.com",
  password: "1234",
  type: "rider",
  drivers_license_number: "1234",
};
async function start() {
  user.id = 1;
  /*
  const res = await register(user);
  const res = await login({
    email: user.email,
    password: user.password,
  });
  */
  const userId = await me();
  console.log(userId);
}

start();
