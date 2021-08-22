# Recipe Stats Calculator

This calculator processes `hf_test_calculation_fixtures.json` file with recipe data to calculate following stats:

1. Count the number of unique recipe names.
2. Count the number of occurrences for each unique recipe name (alphabetically ordered by recipe name).
3. Find the postcode with most delivered recipes.
4. Count the number of deliveries to postcode `10120` that lie within the delivery time between `10AM` and `3PM`, examples _(`12AM` denotes midnight)_:
    - `NO` - `9AM - 2PM`
    - `YES` - `10AM - 2PM`
5. List the recipe names (alphabetically ordered) that contain in their name one of the following words:
    - Potato
    - Veggie
    - Mushroom

## Setup

### Prerequisite
1. go1.16
2. Docker for Desktop 3+ version [More info](https://www.docker.com/)
3. Make sure you have correctly placed `hf_test_calculation_fixtures.json` file in `resources/fixtures/` directory, existing one is stripped version of [original](https://test-golang-recipes.s3-eu-west-1.amazonaws.com/recipe-calculation-test-fixtures/hf_test_calculation_fixtures.tar.gz) file.

### Clone Repository
Clone this repository somewhere on your computer and `cd recipe-stats-calculator`
```
git clone git@github.com:muzfr7/recipe-stats-calculator.git
```

## Usage
Make sure you are at the root of the project and execute following command to build Docker image and execute binary
> This command will basically build the image and run CLI binary with expected output.
> Use this command after making changes to code!
```
make up
```

Use following command to run the binary again
> This command will directly execute binary from Docker container without building the image again
```
make run
```
> above command will produce expected output in following format.
```json
{
    "unique_recipe_count": 15,
    "count_per_recipe": [
        {
            "recipe": "Mediterranean Baked Veggies",
            "count": 1
        },
        {
            "recipe": "Speedy Steak Fajitas",
            "count": 1
        },
        {
            "recipe": "Tex-Mex Tilapia",
            "count": 3
        }
    ],
    "busiest_postcode": {
        "postcode": "10120",
        "delivery_count": 1000
    },
    "count_per_postcode_and_time": {
        "postcode": "10120",
        "from": "11AM",
        "to": "3PM",
        "delivery_count": 500
    },
    "match_by_name": [
        "Mediterranean Baked Veggies", "Speedy Steak Fajitas", "Tex-Mex Tilapia"
    ]
}
```

Use following command to stop and remove app container
```
make down
```

## Testing

Use following command to run all tests
> This command will also generate test coverage file in `testdata/` directory
```
make tests
```

Use following command to view generated test coverage in browser
```
make coverage
```

## Notes

#### Regarding Json file reading

Reader -> Read() method in `app/infrastructure/fs/json/json.go` is used to parse given Json file and return each Json object in reader channel which is then received in `cmd/cli/main.go` and passed into Service -> CalculateStats() method in `app/usecases/recipe/service.go` for stats calculations.

> Reader -> Read() method doesn't load whole Json file in memory, it instead reads each row gradually using json.NewDecoder()

#### Regarding a functional requirement #4 and 5

Following is used in `cmd/cli/main.go` to count the number of deliveries to postcode 10120 that lie within the delivery time between 10AM and 3PM (Functional req #4).
```
filterByPostcodeAndTime := recipeDomain.DeliveriesByPostcodeAndTime{
   Postcode: "10120",
   From:     10,
   To:       3,
}
```

Following is used in `cmd/cli/main.go` to filter the recipe names that contain given words in their name (Functional req #5).
```
filterByWords := []string{"Potato", "Veggie", "Mushroom"}
```
