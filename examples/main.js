import { group, check } from 'k6';
import json from "k6/x/json";

export const options = {
  iterations: 1,
  vus: 1,
};

export default function () {

  group("JSON Marshal", function () {
    const sourceObj = {
      "name": "Xacobe",
      "age": 30,
      "city": "Arteixo",
    };

    const marshalResult = json.marshal(sourceObj); // JSON string
    console.log("Marshal result:", marshalResult);

    check(marshalResult, {
      "Marshal should return a string": (r) => typeof r === "string",
      "The marshalled object should contain the sourceObj properties": (r) => {
        const obj = JSON.parse(r);
        return obj.name === sourceObj.name && obj.age === sourceObj.age && obj.city === sourceObj.city;
      },
    });
  });


  group("JSON Unmarshal", function () {
    const sourceJson = `{"name":"Uxia","age":32,"city":"Santiago de Compostela"}`;

    const unmarshalResult = json.unmarshal(sourceJson); // JSON Object
    console.log("Unmarshal result:", unmarshalResult);

    check(unmarshalResult, {
      "Unmarshal should return an object": (r) => typeof r === "object",
      "The unmarshalled object should contain the sourceJson properties": (r) => {
        return r.name === "Uxia" && r.age === 32 && r.city === "Santiago de Compostela";
      },
    });
  });


  group("Unmarshal failures", function () {
    const invalidJSON = `{"name": "Brais", "age": 30`; // Missing closing bracket
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

    const marshalResult = json.marshal(sourceObj); // empty string

    console.log("Marshal Failure result:", marshalResult);
    check(marshalResult, {
      "Marshal with circular reference should return empty string": (r) => r === "",
    });
  });

}
