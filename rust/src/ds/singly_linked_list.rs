type Link <T> = Option<Box<Node<T>>>;

struct Node<T> {
    val: T,
    next: Link<T>,
}

pub struct Stack<T> {
    head: Link<T>
}

impl <T> Stack<T> {
    pub fn new() -> Self {
        Stack {
            head: None,
        }
    }

    pub fn push(&mut self, val: T) {
        let new_node = Box::new(Node{
            val,
            next: self.head.take() // take() takes the value out of Option<> and leaves a None.
        });

        self.head = Some(new_node)
    }

    pub fn pop(&mut self) -> Result<T, &str> {
        match self.head.take() {
            None => Err("stack is empty"),
            Some(node) => {
                self.head = node.next;
                Ok(node.val)
            }
        }
    }

    pub fn is_empty(&self) -> bool {
        self.head.is_none()
    }
}

impl <T> Default for Stack<T> {
    fn default() -> Self {
        Self::new()
    }
}