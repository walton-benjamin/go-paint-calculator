# go-paint-calculator
Created by Benjamin Walton, for GoLang programming @ TSI




# Functionality
## Walls
- Asks the user to enter a number of walls they wish to paint.
- Asks what format the user would like to enter the wall sizes, either as a total surface area, or as LxW dimentions .
- Iteratively asks for the dimensions of each wall in the room to be painted.


## Paint
- Displays a formatted map of coloured paints, and price per litre of the paints.
- Allows user to select a paint colour, by typing in the colour paint .
    - If a custom colour is selected, allows the user to describe the colour they would like
- Displays an initial price, and number of tins needed to paint the wall.



## Optional Extras
- Displays optional extra a user may wish to purchase at the same time. Paint brushes, overalls etc.
- User can accept or reject adding extras. If they reject, a receipt it processed.
- If they accept, the user can select an extra item using either it's full name (non case sensitive) or using #x notation (where x is an integer in the displayed list).
- The user can also enter 'i/item/items' to display the list of available items again.
- ...or enter 'x/exit/c/cancel' to cancel adding any extra items.
- Regardless of how the user enters the extra item identifer, the program checks if the given identifier is in the available items.
    - if not, will notify the user, and ask them to try again.
- If found, will query how many of the item, the user wants to order.
- until a valid integer input is given, will continue to ask for a valid input.
- when a valid number is given, will add that many of the selected extra to the order slice.
- continues asking if the user wants to add extra items until told no.


## Receipt
- When order is complete, the program processes and formats a new receipt for the user
- First, prints the walls & surface area entered
    - If the user enters surface area by dimensions, program will also return the dimentions entered
- Displays the total surface area entered, 

- Displays the chosen paint colour
    - if custom colour is added, also displays the description of the colour provided.
- Lists the volume of each tin of this colour paint, and the price of each tin.
- States the area the tin lists it will cover, and calculates how many tins will be needed to cover the needed surface area
- Displays a total cost of the tins needed

- If no extras are added, extras will show "None"
- If any extras are added, will display the extra added & the quantity of each item
- Also displays the total amount spent on each set of extras

- Generates a Subotal of costs for paint & extras.
- Calculates VAT ontop of order (static 20%)
- Displays the total for the user to pay
- Closes the receipt and thanks the user


# Non Functional features
- Paints are displayed in line with eachother, despite variable length names.
    - Dynamically extends to the longest name length.
- Error handling
- If zero items are entered, will remove the extra from the list to be added.


# Scope for potential additions
- Functionality to add multiple rooms, with multiple paint colours
- Ability to add a different colour for each wall in a given room
- Functionality to have a variable width receipt (only need to change the hard coded spacers)
- if multiple instances of the same extra are added, could combine them into one row in the receipt