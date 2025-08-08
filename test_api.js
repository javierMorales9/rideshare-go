const url = "http://localhost:8080"
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQ2NzgxODcsInVzZXJfaWQiOjF9.n9LSSWw4yzubadAuCq3s3V4Q9rGjFev5XNN1Zw3wRaQ";

async function makePetition(path, method, headers, body) {
  const res = await fetch(`${url}/${path}`, {
    method,
    headers: {
      Authorization: `Bearer ${token}`,
      ...headers,
    },
    body: body && JSON.stringify(body),
  });

  if (!res.ok) {
    return { error: res.statusText, data: await res.text() }
  }

  return { data: await res.json() };
}

const rider = {
  first_name: "Pedrito",
  last_name: "Lanzadera",
  email: "pedrolanza@gmail.com",
  password: "1234",
  type: "rider",
};
const driver = {
  first_name: "Chapo",
  last_name: "Guzman",
  email: "melasuda@gmail.com",
  password: "1234",
  type: "driver",
  drivers_licens_number: "1234",
}

async function start() {
  rider.id = 1;
  driver.id = 2;
  /*
  const res = await register(driver);
  return makePetition('auth/login', "POST", {
    email: rider.email,
    password: rider.password,
  });
  const riderId = await makePetition('api/me', "GET");
  const res = await makePetition('api/trip_requests', "POST", undefined, {
      rider_id: riderId,
      start_address: "Calle de la piruleta 12 Madrid",
      end_address: "Avenida Gerente Peinado 45 5ºC",
      state: "MA"
  });
  const res = await makePetition(`api/trip_requests/8`, "GET");
  */
  const res = await makePetition(`api/trips`, "GET");
  const res2 = await makePetition(`api/trips/my?rider_id=1`, "GET");

  console.log(res.data, res2.data);
}

start();
