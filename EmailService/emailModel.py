from enum import Enum
from pydantic import BaseModel, EmailStr
from typing import List, Optional

class BodyType(str, Enum):
    plain = "plain"
    html = "html"

class EmailModel(BaseModel):
    recipients: List[EmailStr]
    subject: str
    body: str
    body_type: Optional[str] = BodyType.plain