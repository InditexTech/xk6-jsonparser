import { group, check } from 'k6';
import json from "k6/x/json";

export const options = {
  iterations: 1,
  vus: 1,
};

export default function () {

  group("JSON Marshal & Unmarshal", function () {
    const sourceObj = {
      "name": "John",
      "age": 30,
      "city": "New York",
    };

    const marshalResult = json.marshal(sourceObj); // JSON string
    const unmarshalResult = json.unmarshal(marshalResult); // JavaScript object (equals to `source`)

    console.log("Marshal result:", marshalResult);
    console.log("Unmarshal result:", unmarshalResult);

    check(marshalResult, {
      "Marshal should return a string": (r) => typeof r === "string",
    });

    check(unmarshalResult, {
      "Unmarshal should return a JavaScript object": (r) => typeof r === "object",
      "The unmarshalled object should be equal to the source object": (r) => r.name === sourceObj.name && r.age === sourceObj.age && r.city === sourceObj.city,
    });
  });


  group("Unmarshal failures", function () {
    const invalidJSON = "{name: 'John', age: 30, city: 'New York'}";
    const unmarshalResult = json.unmarshal(invalidJSON); // null

    console.log("Unmarshal Failure result:", unmarshalResult);
    check(unmarshalResult, {
      "Unmarshal with invalid JSON should return null": (r) => r === null,
    });
  });


  group("Marshal failures", function () {
    const sourceObj = {
      "name": "John",
      "age": 30,
      "city": "New York"
    }

    //Add circular reference to create invalid JSON
    sourceObj.circularReference = sourceObj;

    const marshalResult = json.marshal(sourceObj); // null

    console.log("Marshal Failure result:", marshalResult);
    check(marshalResult, {
      "Marshal with circular reference should return empty string": (r) => r === "",
    });
  });

}
