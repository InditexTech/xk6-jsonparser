import { check } from 'k6';
import json from "k6/x/json"
import { expect } from "./k6chaijs.js";

export const options = {
  iterations: 1,
  vus: 1,
};

export default function () {
  let data = '{"key1":"valorParaKey1String", "key2":true, "key3": 1}'
  let objectParsed = json.unmarshal(data)

  check(json.unmarshal(data), {'is valid JSON Key2 Value': (d) => d.key2,});

  //Invalid JSON String
  let badData = '<html>"key1":"valorParaKey1String"</html>'
  let objectBadParsed = json.unmarshal(badData)
  //null is returned
  expect(objectBadParsed).to.equal(null)

  let objectData  = {
    id: "id",
    name: "SRE",
    age: 30,
    email: "test@inditex.com",
    url: "https://www.inditex.com&date=2021-09-01T00:00:00Z"
  };

  let objectJson = json.marshal(objectData)
  let correctJson = `{"age":30,"email":"test@inditex.com","id":"id","name":"SRE","url":"https://www.inditex.com&date=2021-09-01T00:00:00Z"}\n`
  expect(objectJson).to.equal(correctJson)

  //Add circular reference to create invalid JSON
  objectData.circularReference = objectData;
  let objectBagJson = json.marshal(objectData)
  //empty string is returned
  expect(objectBagJson).to.equal(``)
}
