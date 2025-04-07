import http from 'k6/http';
import { check, sleep } from 'k6';

// Configuration options
export let options = {
  vus: 5,           // Number of virtual users
  duration: '5s',    // Duration of the test (5 seconds in this case)
};

export default function () {
  // The API endpoint and the payload for the POST request
  const url = 'http://host.docker.internal:8080/api/test/banner/getlist';
  const payload = JSON.stringify({
    page: 1,
    perPage: 10,
    searchText: '',
    sortBy: {
      field: 'updatedDate',
      mode: 'desc',
    },
  });

  // HTTP headers
  const headers = {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
  };

  // Send the POST request
  let res = http.post(url, payload, { headers: headers });

  // Check the response status
  check(res, {
    'is status 200': (r) => r.status === 200,
  });

  // Sleep to simulate a delay between requests (e.g., 1 second)
  sleep(1);
}


