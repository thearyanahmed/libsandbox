from dotenv import load_dotenv
from pathlib import Path

# enable .env file support
def load_env():
    env_path = Path('') / '.env'
    load_dotenv(dotenv_path=env_path)
