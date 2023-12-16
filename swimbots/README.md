# Swimbots GenePool Simulation

## Description
- This program simulate a population of artificial life form (Swimbots) living in an imaginary pool (GenePool), where they search for food and produce offspring.
- User can observe different dominant evolutionary traits being preserved by tweaking a group of parameters (mating preferecne, amount of food in the pool, maximum age of the Swimbots, etc.) that affects the genepool.
- As each simulation finished, the program will produce a gif file which record the simulation and a report containing graph that shows information about the dominant traits in our program.

## Instructions
- Here we provide the instructions of running the simulation.
    1. Make sure golang was installed in your computer.  You can make check by running the following command in terminal.
         ```sh
        go version
        ```
        (Install golang [here](https://go.dev/doc/install) if it doesn't recognize the command.)
    2. Download GenePool_Package.zip and unzip the file
    3. Move all the folder within into your go/src. (The folder should contain 4 folders: canvas, gifhelper, gogif, swimbots)
    3. Open your terminal and enter the swimbots folder. Build the program by running
         ```sh
        go build
        ```
    4. Run the program by running
        ```sh
        ./swimbots
        ```
    - After the program started you will see a prompt:
        ```
        Welcome to swimbot genepool simulation.
        Would you like to simulate the genepool with default parameters? (y/n)
        ```

- If you would like to run the simulation with defaul parameters, please enter y. Otherwise, please enter n.
    - If you enter y, the simulation will start running. You should see:
        ```
        Simulating genepool with default parameters.
        Number of generations:  1000
        Time interval:  1
        Initial number of bots:  200
        Number of food bits we add every time we add food:  5
        View range of a swimbot:  300
        Proximity for swimbot to eat or mate: 10
        Energy of the foodBits 50
        Energy threshold for hungry:  50
        Maximum age of a bot:  1000
        The frequency of putting in food:  5
        The mass of each segment: 10
        The energy loss factor is: 0.0005
        The mating prefernce is randomized.
        Parameters received. Start Simulation!
        Images drawn!
        Making GIF.
        Animated GIF produced!
        Existing normally.
        ```
        - After the simulation was finished, you should see pond.out.gif, results.txt, and the folder csvFiles in the folder.
    - If you would like to specify the parameters yourself, please enter n.
        - Here is the list of the parameters that you could specify:
            - Number of generations in the simulation (numGen)
            - Time interval between each generation
            - Initial number of Swimbots in the GenePool
            - Number of food bits we add every time we add food
            - View range of a Swimbot (how far can a Swimbot see when it's looking for mate and food)
            - Proximity for swimbot to eat or mate (The maximum distance between a Swimbot and it's target to initiate an action)
            - Energy of the foodBits
            - Energy threshold for Swimbot to become hungry and start to search for food instead of mate
            - Maximum age of a bot
            - The frequency of putting in food (in generation)
            - The mass of each segment
            - The energy loss factor (How fast does the Swimbot loses its energy when it swims)
            - The mating prefernce (The preference of Swimbots when they look for a mate.)
                -  If choose 0, the Swimbots choose its mate randomly
                -  If choose 1, the Swimbots prefer to choose a mate with more segments
                -  If choose 2, the Swimbots prefer to choose a mate with less segments
                -  If choose 3, the Swimbots prefer to choose a mate that swim faster
                -  If choose 4, the Swimbots prefer to choose a mate with similar number of segments
                -  If choose 5, the Swimbots prefer to choose a mate with similar main segment length.
    - The prompt will ask for each parameters from the user. After entering them, the prompt will also generate a list of parameters used in this simulation.
    - When the simulation is finished, you should see the pond.out.gif in the swimbot folder as well. (Note that the previous gif will get overwritten. If you would like to preserve the gif for previous simulations, please change the name of the gif before you run the next simulation.)

## Analysis
- The simulation generated some written analysis of the results in the same folder, under the Results.txt file.
- There are also csv files generated that should be exported to python to visualize the graphs.

- Here we provide the instructions of generating the graphs and results.
    - If you would like to generate a PDF file which present the result of the simulation, we provide two methods
        1. Go to the GoogleColab notebook we authored at [here](https://colab.research.google.com/drive/16-ietCQeOcxhZpm05NCFrXazz-7XP-rR?usp=sharing)
        2. Make sure you connect to a runtime.
        3. Go to the folder icon, click on "Upload to session storage", and select all the files generated in the csvFiles folder, as well as the results.txt file, and upload them into your current session's storage.
        4. Go to Runtime-> Run All to run the whole notebook and generate the the graphs and results in a pdf. The results will be saved in a Results folder, that has one folder with the graphs (graphs_generated folder) and the pdf of the results combined with the graphs in Analysis.pdf.
        5. Make sure to download them because uploaded and generated files will get deleted when this runtime is restarted.
    - If you would like to generate PDF files locally:
        1. Make sure you have the newest version of python. (In stall python [here](https://www.python.org/downloads/), if it's not installed)
        2. After downloading python, you should have pip in your package. Install the required packages by running the following commands:
            ```
            pip install pandas
            pip install matplotlib
            pip install numpy
            pip install fpdf
            ```
        3. After all the package is installed, run the following command:
            ```
            python3 swimbotsanalysis.py
            ```
        4. You should see a file "Analysis.pdf" and a folder called "Graphs_Generated" which contained all the graphs being generated in the folder "Results".
