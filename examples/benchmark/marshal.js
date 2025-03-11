import { group } from 'k6';
import json from "k6/x/json";

const native = __ENV.NATIVE === "1";

export default function () {
  const sourceObj = {
    "user": {
      "id": "ajf92jf9",
      "first_name": "xkdaslqw",
      "last_name": "pqoweiur",
      "address": {
        "street": "asdklj123",
        "city": "qowieur",
        "postal_code": "zxcmvn123"
      }
    },
    "order": {
      "order_id": "zxcvbn12",
      "date": "klasjdf093",
      "products": [
        {
          "product_id": "lkjasd812",
          "name": "zmxncviqwe",
          "price": "098asdjf"
        },
        {
          "product_id": "qwerlkjas",
          "name": "zxcmnbv912",
          "price": "234lkjasd"
        }
      ]
    },
    "transaction": {
      "transaction_id": "lkajsdf0923",
      "payment_method": "zmxncvoiqw",
      "status": "alksdjf1209"
    }
  };

  if (native) {
    group("JSON Marshal (native)", function () {
      JSON.stringify(sourceObj);
    });
  } else {
    group("JSON Marshal (extension)", function() {
      json.marshal(sourceObj);
    });
  }

}
