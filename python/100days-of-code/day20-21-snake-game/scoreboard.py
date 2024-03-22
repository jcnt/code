from turtle import Turtle

FONT = ("Arial", 24, "normal")

class ScoreBoard(Turtle):

    def __init__(self):
        super().__init__()
        self.score = 0
        with open("highscore.txt") as file:
            self.highscore = int(file.read())
        self.color("white")
        self.penup()
        self.goto(0, 270)
        self.hideturtle()
        self.update_scoreboard()

    def update_scoreboard(self):
        self.clear()
        self.write(f"Score: {self.score}, High score: {self.highscore}", align="center", font= FONT )

#    def game_over(self):
#        self.color("white")
#        self.penup()
#        self.goto(0, 0)
#        self.write("GAME OVER", align = "center", font = FONT)

    def reset(self):
        if self.score > self.highscore:
            self.highscore = self.score
            with open("highscore.txt", mode="w") as file:
                file.write(f"{self.highscore}")
        self.score = 0
        self.update_scoreboard()


    def increase_score(self):
        self.score += 1
        self.clear()
        self.update_scoreboard()
