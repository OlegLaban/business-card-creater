# Business card creator

This tool you can use for making pdf file witch you can print as weel and images will be in rigth position for both sides.

## For example
You have two images first_side.png and second_side.png. 
If you run this command:
./main -f "first_side.png" -w 40 -s "second_side.png" -p 40 -b 1 -j 5 -k 5 -z 5 -l 5

You will receive a PDF file with two pages. On the first page, you will see "first_size.png" with the following properties:

    Width: 40px

    Border: 1px

    Margin from the left edge (border): 5px

    Margin from the top edge (border): 5px

    Offset from the top of the page: 5px

    Offset from the left of the page: 5px

On the second page, you will see "second_side.png" with the same properties. However, its position will be mirrored compared to the first page.
