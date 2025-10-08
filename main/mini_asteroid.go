components {
  id: "mini_asteroid"
  component: "/main/mini_asteroid.script"
}
components {
  id: "mineable"
  component: "/main/mineable.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"asteroid0\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/game.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.02
    y: 0.02
    z: 0.02
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"asteroids\"\n"
  "mask: \"enemies\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 12.500291\n"
  "}\n"
  ""
}
embedded_components {
  id: "factory"
  type: "factory"
  data: "prototype: \"/main/unobtainium.go\"\n"
  "load_dynamically: true\n"
  ""
}
