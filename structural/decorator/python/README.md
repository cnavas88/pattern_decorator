## How to run python project

Steps:
1. Build docker image: 

```bash
docker build -t NAME_IMAGE .
```

2. Run docker image: 

```bash   
docker run NAME_IMAGE
```

### Execute the tests

The project is configured to run tests automatically when you run docker image. See dockerfile.

### app.py?

if you want execute an example of project. YOu can create app.py file and copy the next content:

```python
import coffeeshop

# First we generate a concrete coffee.
myCoffee = coffeeshop.Concrete_Coffee()

print('> Concrete Coffee')
print('Ingredients: '   + myCoffee.get_ingredients())  
print('Cost: '          + str(myCoffee.get_cost())) 
print('Sales tax: '     + str(myCoffee.get_tax()))

# Adding milk to my coffee
myCoffee = coffeeshop.Milk(myCoffee)
print('\n>> Concrete Coffee with milk')
print('Ingredients: '   + myCoffee.get_ingredients())  
print('Cost: '          + str(myCoffee.get_cost())) 
print('Sales tax: '     + str(myCoffee.get_tax()))

# Adding sugar to my coffee
myCoffee = coffeeshop.Sugar(myCoffee)
print('\n>> Concrete Coffee with milk and sugar')
print('Ingredients: '   + myCoffee.get_ingredients())  
print('Cost: '          + str(myCoffee.get_cost())) 
print('Sales tax: '     + str(myCoffee.get_tax()))

# Adding vanilla to my coffee
myCoffee = coffeeshop.Vanilla(myCoffee)
print('\n>> Now I want a vanilla in my coffee:')
print('Ingredients: '   + myCoffee.get_ingredients())  
print('Cost: '          + str(myCoffee.get_cost())) 
print('Sales tax: '     + str(myCoffee.get_tax()))
```

In dockerfile you can uncomment line that execute app.py.