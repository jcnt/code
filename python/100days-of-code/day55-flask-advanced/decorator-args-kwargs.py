## Advanced Python Decorator Functions

class User: 
    def __init__(self, name) -> None:
        self.name = name
        self.is_logged_in = False

def is_authenticated_decorator(function): 
    def wrapper(*args, **kwargs):
        if args[0].is_logged_in == True:
            function()
    return wrapper

def create_blog_post(user):
    print(f"This is {user.name}'s new blog post.")

new_user = User("bela")
create_blog_post(new_user)

